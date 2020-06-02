import React from 'react'

const Header = () => {
    return (
        <React.Fragment>
            <div className="rowheader">
                    {/* <div className="column">
                        <div className="blue-column">
                            Transaction Id
                      </div>
                    </div> */}
                    <div className="column">
                        <div className="green-column">
                            Merchant ID
                      </div>
                    </div>
                    <div className="column">
                        <div className="blue-column">
                            Terminal ID
                      </div>
                    </div>
                     <div className="column">
                        <div className="green-column">
                            Card Number
                      </div>
                   </div>
                    <div className="column">
                        <div className="blue-column">
                            Amount
                      </div>
                    </div>
                   <div className="column">
                        <div className="green-column">
                            Currency
                      </div>
                    </div>
                   {/* <div className="column">
                        <div className="blue-column">
                            Trx Date
                      </div>
                    </div> */}
                    <div className="column">
                        <div className="green-column">
                            AF Amount
                      </div>
                    </div>
                    <div className="column">
                        <div className="blue-column">
                            AF Rate
                      </div>
                    </div>
                   {/*    <div className="column">
                        <div className="green-column">
                            Card Presnent
                      </div>
                    </div> */}
                    
                </div>
                {/* <hr style={{height: "10px", visibility:"Hidden"}} />  */}
        </React.Fragment>
    )
}

export default Header
