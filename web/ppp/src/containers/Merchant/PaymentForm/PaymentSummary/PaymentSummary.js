import React from 'react'

const PaymentSummary = (props) => {

    if(props.show)
    {
        return (
        <React.Fragment>
            <nav className="panel-left">
                <div>Form No = {props.data}</div>           
            </nav>
        </React.Fragment>
        )
    }
    return null
}

export default PaymentSummary
