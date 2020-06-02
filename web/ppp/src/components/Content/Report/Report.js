import React from 'react'
import { withRouter } from 'react-router-dom'




const Report = (props) => {

    return (
        <React.Fragment>
            <div className="rowContent">
                <div className="row" onMouseOver={props.showSummaryOnHover} onMouseOut={props.hideSummaryOnHover}>
                    <div className="column"><div className="blue-column">{props.merchantname}</div></div>
                    <div className="column"><div className="blue-column">{props.merchantid}</div></div>
                    <div className="column"><div className="blue-column">
                        <a href={"file:" + props.pdffile} download>
                            <img src={props.csv} alt="csv format" width="20px" height="20px" />
                        </a>
                    </div></div>
                    <div className="column"><div className="blue-column">
                        <a href={props.xlsfile} download>
                            <img src={props.pdf} alt="pdf format" width="20px" height="20px" />
                        </a>
                    </div></div>
                    <div className="column"><div className="blue-column">
                        <a href={props.csvfile} download>
                            <img src={props.xls} alt="xls format" width="20px" height="20px" />
                        </a>
                    </div></div>
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

export default withRouter(Report)
