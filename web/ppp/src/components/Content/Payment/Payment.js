import React from 'react'
import { withRouter } from 'react-router-dom'



const Payment = (props) => {
    return (
        <React.Fragment>
            <div className="rowContent">
                <div className="row">
                    {/* <div className="column"><div className="blue-column">{props.transactionid}</div></div> */}
                    <div className="column"><div className="blue-column ">{props.merchantid}</div></div>
                   <div className="column"><div className="blue-column">{props.terminalid}</div></div>
                     <div className="column"><div className="blue-column">{props.cardnumbermasked.replace(/[0-9](?=([0-9]{4}))/g, '*')}</div></div>
                    <div className="column"><div className="blue-column">{props.originaltransactionamount}</div></div>
                    <div className="column"><div className="blue-column">{props.foreigncurrency}</div></div>
                    {/* <div className="column"><div className="blue-column">{props.transactiondate}</div></div> */}
                    <div className="column"><div className="blue-column">{props.originaltransactionamount * 0.1}</div></div>
                     <div className="column"><div className="blue-column">{props.marginrate}</div></div>
                   {/* <div className="column"><div className="blue-column">{props.afcardpresence}</div></div>
                    <div className="column"><div className="blue-column">{props.transactionid}</div></div>
                    <div className="column"><div className="blue-column">{props.merchantid}</div></div>
                    <div className="column"><div className="blue-column">{props.terminalid}</div></div>
                    <div className="column"><div className="blue-column">{props.originaltransactionamount}</div></div>
                    <div className="column"><div className="blue-column">{props.foreigncurrency}</div></div>
                    <div className="column"><div className="blue-column">{props.transactiondate}</div></div>
                    <div className="column"><div className="blue-column">{props.acquirerfee}</div></div>
                    <div className="column"><div className="blue-column">{props.marginrate}</div></div>
                    <div className="column"><div className="blue-column">{props.afcardpresence}</div></div> */}
                     
                </div>
            </div>

        </React.Fragment >

    )
}

export default withRouter(Payment)


