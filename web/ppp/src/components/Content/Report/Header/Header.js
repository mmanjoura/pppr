import React from 'react'

const Header = (props) => {
    return (
        <React.Fragment>

            <div className="rowheader">
                <div className="column">
                    <div className="blue-column">
                        Merchant Name
                      </div>
                </div>
                <div className="column">
                    <div className="green-column">
                        Merchant ID
                      </div>
                </div>
                <div className="column">
                    <div className="green-column">
                        PDF File
                      </div>
                </div>
                <div className="column">
                    <div className="blue-column">
                        Xsl File
                      </div>
                </div>
                <div className="column">
                    <div className="green-column">
                        Csv File
                      </div>
                </div>
                <div className="column">
                    <div className="green-column">
                        Approved
                      </div>
                </div>
            </div>
            {/* <hr style={{height: "10px", visibility:"Hidden"}} />  */}
        </React.Fragment>
    )
}

export default Header
