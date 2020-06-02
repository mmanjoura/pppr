import React, { Component } from 'react';
import GrpPayment from '../../../components/Content/GrpPayment/GrpPayment';
import Header from '../../../components/Content/GrpPayment/Header/Header';
import Approved from './../../../assets/images/approved.png'

export default class Payments extends Component {

    state = {
        acquirer: '',
        numPayments: '',
        transactions: '',
        amount: '',
        fees: '',
        viewPaymentDetails: false,
        showSummary: false,
        payments: [],
        paymentTrxs: [],
        formType: 2,
        merchantId: '',
    }

    viewPaymentDetailsHandler = (id) => {
        let midTrxs = this.props.payments.filter(p => {
            return p.merchantid === id;
        });
        this.setState({
            viewPaymentDetails: true,
            merchantId: id
        }, () => {
            this.loadData(id);
        });


    }

    loadData = async () => {
        const { merchantId } = this.state;
        console.log("Merchant ---------", merchantId);

        // Get data via XHR...

        this.props.paymentDetailsHandler(merchantId)
    }

    showSummaryOnHoverHandler = (id) => {
        let midTrxs = this.props.payments.filter(p => {
            return p.merchantid === id;
        });
        this.setState({
            showSummary: false,
            paymentTrxs: midTrxs

        });

    }

    render() {

        // Group by Mids
        const groupBy = (list, keyGetter) => {
            const map = new Map();
            list.forEach((item) => {
                const key = keyGetter(item);
                const collection = map.get(key);
                if (!collection) {
                    map.set(key, [item]);
                } else {
                    collection.push(item);
                }
            });
            return map;
        }

        // Group Payment by MIDs
        let pytsMap = groupBy(this.props.payments, p => p.merchantid);
        // console.log("pyts", pytsMap)

        // Convert Map to Array
        let MapToArrayPayments = [...pytsMap]

        // Iterate through the map
        const gPayments = MapToArrayPayments.map((p, index) => {
            const groupedMidsPayments = p.map((val, index) => {
                return val
            });

            let obj = groupedMidsPayments[1]

            // Now build your State, the Sum of Amounts
            const sumTransactionAmounts = (groupedMidsPayments) => {
                const summed = groupedMidsPayments[1].reduce((acc, current) => {
                    const key = current.originaltransactionamount;

                    // Retreive the previous price from the accumulator
                    const previousAmount = acc[key]; // Might also return undefined

                    // Create your temp current price value, and be sure to deal with numbers.
                    let currentAmount = Number(current.originaltransactionamount);

                    // If you had a previous value (and not undefined)
                    if (previousAmount) {
                        // Add it to our value
                        currentAmount += Number(previousAmount);
                    }
                    // Return the future accumulator value
                    return Object.assign(acc, {
                        [key]: currentAmount, // new values will overwrite same old values
                    })
                }, {})

                // Return an array of each value from the summed object to our sortedArray
                const sortedArray = Object.keys(summed).sort().map((val) => {
                    return summed[val];
                });

                // Sum the values
                const sum = sortedArray.reduce((a, b) => {
                    return a + b;
                }, 0);

                return sum;
            };

            // Now build your State, the Sum of Amounts
            const sumAcquirerFee = (groupedMidsPayments, index) => {
                const summed = groupedMidsPayments[1].reduce((acc, current) => {
                    const key = current.originaltransactionamount;

                    // Retreive the previous price from the accumulator
                    const previousAmount = acc[key]; // Might also return undefined

                    // Create your temp current price value, and be sure to deal with numbers.
                    let currentAmount = Number(current.fee);

                    // If you had a previous value (and not undefined)
                    if (previousAmount) {
                        // Add it to our value
                        currentAmount += Number(previousAmount);
                    }
                    // Return the future accumulator value
                    return Object.assign(acc, {
                        [key]: currentAmount, // new values will overwrite same old values
                    })
                }, {})

                // Return an array of each value from the summed object to our sortedArray
                const sortedArray = Object.keys(summed).sort().map((val) => {
                    return summed[val];
                });

                // Sum the values
                const sum = sortedArray.reduce((a, b) => {
                    return a + b;
                }, 0);

                return sum;
            };

            return (
                <React.Fragment key={obj[0].transactionid}>
                    <GrpPayment

                        merchantname={obj[0].merchantname}
                        numPayments={groupedMidsPayments[1].length}
                        numtransactions={groupedMidsPayments[1].length}
                        totalamount={sumTransactionAmounts(groupedMidsPayments)}
                        acquirerfee={sumAcquirerFee(groupedMidsPayments)}

                        // P[0] is the MID viewDetailsclicked
                        viewDetailsclicked={() => this.viewPaymentDetailsHandler(p[0])}
                        show={this.state.showSummary}
                        approved={Approved}

                    />
                </React.Fragment>

            )

        });

        return (
            <React.Fragment>
                <br />
                <Header />
                {gPayments}
            </React.Fragment>

        )
    }
}


