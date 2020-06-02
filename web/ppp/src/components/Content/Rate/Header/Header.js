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
                        Currency
                      </div>
                </div>
                <div className="column">
                    <div className="blue-column">
                        Rate
                      </div>
                </div>                
            </div>
            {/* <hr style={{height: "10px", visibility:"Hidden"}} />  */}
        </React.Fragment>
    )
}

export default Header
