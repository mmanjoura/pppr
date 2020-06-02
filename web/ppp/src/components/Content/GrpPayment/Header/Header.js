import React from 'react'

const Header = () => {
    return (
        <React.Fragment>
            <div className="rowheader">
                    <div className="column">
                        <div className="blue-column">
                            Merchant
                      </div>
                    </div>
                    {/* <div className="column">
                        <div className="green-column">
                            Payments
                      </div>
                    </div> */}
                    <div className="column">
                        <div className="blue-column">
                            Transactions
                      </div>
                    </div>
                    <div className="column">
                        <div className="green-column">
                            Transactions Amount
                      </div>
                    </div>
                     <div className="column">
                        <div className="green-column">
                            AF Amount
                      </div>
                    </div>
                    <div className="column">
                    <div className="green-column">
                        Approved
                      </div>
                </div>
            </div>
          
        </React.Fragment>
    )
}

export default Header
