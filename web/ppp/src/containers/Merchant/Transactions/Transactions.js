import React, { Component } from 'react'
import axios from 'axios'
import PaymentHeader from '../../../components/Content/Payment/Header/Header'
import Payment from '../../../components/Content/Payment/Payment'

export default class Transactions extends Component {

    render() {
        const paymenTrxs = this.props.transactions.map(p => {

            return <Payment
                key={p.paymentid}
                // transactionid={p.transactionid}
                merchantid={p.merchantid}
                terminalid={p.terminalid}
                cardnumbermasked={p.cardnumbermasked}
                originaltransactionamount={p.originaltransactionamount}
                foreigncurrency={p.foreigncurrency}
                transactiondate={p.transactiondate}
                acquirerfee={p.acquirerfee}
                marginrate={p.marginrate}
                afcardpresence={p.afcardpresence}

                IF={p.ifamount}
                SF={p.fcurrencycode}
                generate={null}
                viewTransaction={null}
            />
        });
        return (
            <React.Fragment>
                <PaymentHeader />
                <div id="scroller">
                    {paymenTrxs}
                </div>
            </React.Fragment>
        )
    }
}
