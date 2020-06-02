import React, { Component } from 'react'
import axios from 'axios'
import Header from '../../components/Content/Rate/Header/Header'
import Rate from '../../components/Content/Rate/Rate'
import RateForm from './RateForm/RateForm'
import { ModalSpinner, Modal } from '../../components/Layout/Modal/Modal';

export default class ExchangeRates extends Component {
    constructor(props) {
        super(props)
        this.ratesHandler = this.ratesHandler.bind(this)
        this.state = {
            exchangeRate: {
                createdDate: '',
                createdTime: '',
                rates: [],

            },
            currency: '',
            startDate: null,
            endDate: null,
            showSummary: false,
            loading: false,
            showHeader: false,
        }
    }

    componentDidMount() {
        var todayDate = new Date().toISOString().slice(0, 10);
        axios.get('http://localhost:8040/rates/' + todayDate)
            .then(response => {

                this.setState({
                    exchangeRate: {
                        createdDate: response.data.createddate,
                        createdTime: response.data.createdtime,
                        rates: response.data.rates,
                    }
                });
                 console.log(this.state.exchangeRate)
            });
    }

    showSummaryOnHoverHandler(r) {
        this.setState({
            rate: r,
            showSummary: false

        })
    }

    hideSummaryOnHoverHandler() {
        this.setState({ showSummary: false })

    }

    ratesHandler(currency, start, end) {
       
        this.setState({
            currency: currency,
            startDate: start,
            endDate: end,
        }, () => {
            if (this.state.startDate != null) {
                 var todayDate = new Date().toISOString().slice(0, 10);
                axios.get('http://localhost:8040/rates/' + this.state.startDate.toISOString().slice(0, 10))
                    .then(response => {

                        this.setState({
                            exchangeRate: {
                                createdDate: this.state.exchangeRate.createddate,
                                createdTime: this.state.exchangeRate.createdtime,
                                rates: this.state.exchangeRate.rates.filter(word => word.currency == currency),
                            }
                        });

                    });
            }
        });
    }

    render() {

        if (this.state.loading) {
            return (
                <React.Fragment>
                    <RateForm paymentHandler={this.ratesHandler} />
                    <ModalSpinner />
                </React.Fragment>
            )
        }
        else {
            console.log(this.state.createdTime)
            let exchangeRates = this.state.exchangeRate.rates.map((r, index) => {
                return <Rate
                    key={index}
                    createdDate={this.state.exchangeRate.createdDate}
                    createdTime={this.state.exchangeRate.createdTime}
                    currency={r.currency}
                    rate={r.rate}
                    showSummaryOnHover={() => this.showSummaryOnHoverHandler(r)}
                    hideSummaryOnHover={() => this.hideSummaryOnHoverHandler()}
                />
            });
            return (
                <React.Fragment>
                    <RateForm ratesHandler={this.ratesHandler} />
                    <br></br>
                    {this.state.showHeader == false ? <Header /> : null}

                    <div id="scroller">
                        {exchangeRates}
                    </div>

                </React.Fragment>

            )
        }




    }
}

