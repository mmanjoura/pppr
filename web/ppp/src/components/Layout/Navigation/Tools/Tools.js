import React from 'react'
import { Link } from 'react-router-dom'

const Tools = (props) => {
    return (
        <React.Fragment>
            <div className="profile">
                <div className="icon_wrap">
                    <i className="fas fa-align-justify" />
                </div>
                <div className="profile_dd">
                    <ul className="profile_ul">
                        <li><Link className="" to="/charts"><span className="picon"><i className="" /></span>DASHBOARD</Link></li>
                        <li><Link className="" to="/payments"><span className="picon"><i className="" /></span>PAYMENTS</Link></li>
                         <li><Link className="" to="/payments"><span className="picon"><i className="" /></span>REBATES</Link></li>
                        <li><Link className="" to="/reports"><span className="picon"><i className="" /></span>REPORTS</Link></li>                      
                         <li><Link className="" to="/rates"><span className="picon"><i className="" /></span>RATES</Link></li>
                        <li><Link className="address" to="/logs"><span className="picon"><i className="" /></span>LOGS</Link></li>
                        <li><Link className="settings" to="#"><span className="picon"><i className="" /></span>Settings</Link></li>
                    </ul>
                </div>
            </div>
        </React.Fragment>
    )
}

export default Tools
