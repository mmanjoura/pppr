import React, { Component } from 'react';

import PaymentForm from './PaymentForm/PaymentForm';
import PaymentSummary from './PaymentForm/PaymentSummary/PaymentSummary'
import Transactions from './Transactions/Transactions';
import Payments from './Payments/Payments';
import axios from 'axios';
import { ModalSpinner, Modal } from '../../components/Layout/Modal/Modal';



export default class Merchant extends Component {

    constructor(props) {
        super(props)
        this.paymentHandler = this.paymentHandler.bind(this)
        this.paymentDetailsHandler = this.paymentDetailsHandler.bind(this)
        this.state = {
            viewPaymentDetails: false,
            acquirerId: '',
            merchantId: '',
            startDate: '',
            endDate: '',
            payments: [],
            filteredPayments: [],
            transactions: [],
            loading: false,
        }
    }

    // Set the AcquirerId and Date in payment CMP
    paymentHandler(acquirerId, startDate, endDate) {

        this.setState({
            acquirerId: acquirerId,
            startDate: startDate,
            endDate: endDate,
            loading: true,
        }, () => {

            axios.get('http://localhost:8040/acquirer/payments/' + acquirerId)
                .then(response => {
                    this.setState({
                        payments: response.data,
                    });
                    console.log("Select AcquirerId is: ", acquirerId)
                    // Acquirer Drop down changed
                    const result = this.state.payments.filter(word => word.acquirerid == acquirerId)
                    this.setState({ filteredPayments: result, loading: true }, () => {
                        this.setState({ loading: false })
                    });

                });
        });
    }

    // Show payment Details
    paymentDetailsHandler(mid) {
        this.setState({
            viewPaymentDetails: true,
            merchantId: mid,
        }, () => {
            axios.get('http://localhost:8040/merchant/payments/' + mid)
                .then(response => {
                    this.setState({
                        transactions: response.data,
                        loading: false,
                    });
                    console.log("transactions", this.state.transactions)
                });
        });
    }

    // Close Modal Form when backdrop is clicked
    closeModalHandler = () => {
        this.setState({ viewPaymentDetails: false });

    }

    render() {

        if (this.state.loading) {
            return (
                <React.Fragment>
                    <PaymentForm paymentHandler={this.paymentHandler} />
                    <ModalSpinner />
                </React.Fragment>
            )
        } else {

            return (
                <React.Fragment>
                    <PaymentForm paymentHandler={this.paymentHandler} />
                    <Payments payments={this.state.filteredPayments} paymentDetailsHandler={this.paymentDetailsHandler} />
                    <Modal show={this.state.viewPaymentDetails} modalClosed={this.closeModalHandler}>
                        <Transactions transactions={this.state.transactions} />
                    </Modal>
                </React.Fragment>
            )
        }
    }
}
