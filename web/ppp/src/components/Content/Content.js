import React from 'react'
import { Route } from 'react-router-dom'
import MerchantPayments from '../../containers/MerchantPayments/MerchantPayments'
import MerchantReports from '../../containers/MerchantReports/MerchantReports'
import LogMessages from '../../containers/LogMessages/LogMessages'
import ExchangeRates from '../../containers/ExchangeRates/ExchangeRates'
import BrushBarChart from '../Charts/BrushBarChart/BrushBarChart'
import Merchant from '../../containers/Merchant/Merchant'

const Content = (props) => {
  return (
    <React.Fragment>
      {/* <br></br>
      <Route path="/payments" component={Merchant} />
      {/* <Route path="/paymentdetails" component={MerchantPayments} /> */}
      {/* <Route path="/reports" component={MerchantReports} />
      <Route path="/logs" component={LogMessages} />
      <Route path="/rates" component={ExchangeRates} />
      <Route path="/brushrarchart" component={BrushBarChart} />
      <Route path="/brushrarchart" component={BrushBarChart} /> */} */}
    </React.Fragment>
  )
}

export default Content
