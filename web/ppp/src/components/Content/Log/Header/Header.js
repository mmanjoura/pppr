import React from 'react'

const Header = (props) => {
    return (
        <React.Fragment>
             <div className="rowheader">
                <div className="column">
                    <div className="blue-column">
                        Date
                      </div>
                </div>
                <div className="column">
                    <div className="blue-column">
                        Time
                      </div>
                </div>
                <div className="column">
                    <div className="green-column">
                        Level
                      </div>
                </div>
                <div className="column">
                    <div className="blue-column">
                        Service Name
                      </div>
                </div>
                <div className="column">
                    <div className="green-column">
                        Calling Method
                      </div>
                </div>
                <div className="column">
                    <div className="blue-column">
                        Server Name
                      </div>
                </div>
                <div className="column">
                    <div className="green-column">
                        Latency
                      </div>
                </div>
                <div className="column">
                    <div className="green-column">
                        Message
                      </div>
                </div>
            </div>
            {/* <hr style={{height: "10px", visibility:"Hidden"}} />  */}
        </React.Fragment>
    )
}

export default Header
