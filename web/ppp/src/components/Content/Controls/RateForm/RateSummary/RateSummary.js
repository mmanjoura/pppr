import React from 'react'

const RateSummary = (props) => {

    if(props.show)
    {
        return (
        <React.Fragment>
            <nav className="panel-right">
                <div>1 Euro = {props.rate + " " + props.currency}</div>           
            </nav>
        </React.Fragment>
        )
    }
    return null
}

export default RateSummary
