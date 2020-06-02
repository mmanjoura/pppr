import React from 'react'

const LogSummary = (props) => {

    if(props.show)
    {
        return (
        <React.Fragment>
            <nav className="panel-right">
            <div>Form No = {props.data}</div>           
            </nav>
        </React.Fragment>
        )
    }
    return null
}

export default LogSummary
