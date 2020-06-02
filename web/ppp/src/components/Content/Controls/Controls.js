import React from 'react'
import RateSummary from './RateForm/RateSummary/RateSummary'
import RateForm from './RateForm/RateForm'
import GrpPaymentForm from './GrpPaymentForm/GrpPaymentForm'
import GrpSummary from './GrpPaymentForm/GrpSummary/GrpSummary'
import LogForm from './LogForm/LogForm'
import LogSummary from './LogForm/LogSummary/LogSummary'
import ReportForm from './ReportForm/ReportForm'
import ReportSummary from './ReportForm/ReportSummary/ReportSummary'
import ChartForm from './ChartForm/ChartForm'
import ChartSummary from './ChartForm/ChartSummary/ChartSummary'

const Controls = (props) => {

    switch (props.formType) {
        case 1:
            return (
                <React.Fragment>
                    <RateForm />
                    <RateSummary show={props.show} rate={props.rate.rate} currency={props.rate.currency} />
                    <br />
                </React.Fragment>
            )

        case 2:
        return (
            
            <React.Fragment>
                <GrpPaymentForm />
                <GrpSummary show={props.show} data={props.data}/>
                <br />
            </React.Fragment>
        )

        case 3:
        return (
            
            <React.Fragment>
                <LogForm />
                <LogSummary show={props.show} data={props.data}/>
                <br />
            </React.Fragment>
        )
        case 4:
        return (
            
            <React.Fragment>
                <ReportForm />
                <ReportSummary show={props.show} data={props.data}/>
                <br />
            </React.Fragment>
        )
        case 5:
        return (
            
            <React.Fragment>
                <ChartForm />
                <ChartSummary show={props.show} data={props.data}/>
                <br />
            </React.Fragment>
        )

        default:
        // code block
        return null
    }

}

export default Controls
