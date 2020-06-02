import React from 'react'
import { withRouter } from 'react-router-dom'


const Log = (props) => {
    return (
        <React.Fragment>
            <div className="rowContent">
                <div className="row" onMouseOver={props.showSummaryOnHover} onMouseOut={props.hideSummaryOnHover}>
                    <div className="column"><div className="blue-column">{props.createddate}</div></div>
                    <div className="column"><div className="blue-column">{props.createdtime}</div></div>
                    <div className="column"><div className="blue-column">{props.level}</div></div>
                    <div className="column"><div className="blue-column">{props.servicename}</div></div>
                    <div className="column"><div className="blue-column">{props.callingmethod}</div></div>
                    <div className="column"><div className="blue-column">{props.host}</div></div>
                    <div className="column"><div className="blue-column">{props.latency}</div></div>
                    <div className="column"><div className="blue-column">{props.message}</div></div>
                </div>
            </div>


        </React.Fragment>
    )
}

export default withRouter(Log)
