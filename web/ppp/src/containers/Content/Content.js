import React, { Component } from 'react'
import { Route } from 'react-router-dom'

import Payments from '../Merchant/Payments/Payments'

import Merchant from '../../containers/Merchant/Merchant'

import MerchantReports from '../MerchantReports/MerchantReports'
import LogMessages from '../LogMessages/LogMessages'
import ExchangeRates from '../ExchangeRates/ExchangeRates'

import Charts from '../Charts/Charts';

export default class Content extends Component {
    render() {
        return (
            <React.Fragment>
                <Route path="/payments" component={Merchant} />
                <Route path="/reports" component={MerchantReports} />
                <Route path="/logs" component={LogMessages} />
                <Route path="/rates" component={ExchangeRates} />
                <Route path="/charts" component={Charts} />
            </React.Fragment>
        )
    }
}
