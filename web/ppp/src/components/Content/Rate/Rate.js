import React from 'react'
import { withRouter } from 'react-router-dom'



const Rate = (props) => {
  return (
    <React.Fragment>
      <div className="rowContent">
        <div className="row" onMouseOver={props.showSummaryOnHover} onMouseOut={props.hideSummaryOnHover}>
          <div className="column"><div className="blue-column">{props.createdDate}</div></div>
          <div className="column"><div className="blue-column">{props.createdTime}</div></div>
          <div className="column"><div className="blue-column">{props.currency}</div></div>
          <div className="column"><div className="blue-column">{props.rate}</div></div>
        </div>
      </div>

    </React.Fragment>
  )
}

export default withRouter(Rate)
