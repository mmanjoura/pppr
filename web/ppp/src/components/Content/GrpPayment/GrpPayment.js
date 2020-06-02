import React from 'react'
import { withRouter } from 'react-router-dom'



const GrpPayment = (props) => {
    return (
        <React.Fragment>
            <div className="rowContent">
                <div className="row" onMouseOver={props.showSummaryOnHover} onClick={props.viewDetailsclicked}>
                    <div className="column"><div className="blue-column link">{props.merchantname}</div></div>
                    {/* <div className="column"><div className="blue-column">{props.numPayments}</div></div> */}
                    <div className="column"><div className="blue-column">{props.numtransactions}</div></div>
                    <div className="column"><div className="blue-column">{props.totalamount}</div></div>
                    <div className="column"><div className="blue-column">{props.acquirerfee}</div></div>
                    <div className="column"><div className="blue-column">
                        <a href={props.csvfile} download>
                            <img src={props.approved} alt="approved" width="20px" height="20px" />
                        </a>
                    </div></div>
                </div>
            </div>

        </React.Fragment>
    )
}

export default withRouter(GrpPayment)
