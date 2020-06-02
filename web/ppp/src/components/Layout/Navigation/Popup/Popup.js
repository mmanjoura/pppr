import React, { Component } from 'react'
import axios from 'axios';

export default class Popup extends Component {
   

    render() {
        return (
           <React.Fragment>
            {/* Popup */}
            <div className="popup">
                <div className="shadow" />
                <div className="inner_popup">
                    <div className="notification_dd">
                        <ul className="notification_ul">
                            <li className="title">
                                <p>All Notifications</p>
                                <p className="close"><i className="fas fa-times" aria-hidden="true" /></p>
                            </li>
                            <li className="starbucks success">
                                <div className="notify_icon">
                                    <span className="icon" />
                                </div>
                                <div className="notify_data">
                                    <div className="title">
                                        Planet POS PSW
                                    </div>
                                    <div className="sub_title">
                                    Approved by Colm Sullivan
                                    </div>
                                </div>
                                <div className="notify_status">
                                    <p>Approved</p>
                                </div>
                            </li>
                            <li className="baskin_robbins failed">
                                <div className="notify_icon">
                                    <span className="icon" />
                                </div>
                                <div className="notify_data">
                                    <div className="title">
                                        FIPS-POS.
                                    </div>
                                    <div className="sub_title">
                                        Approval Needed
                                    </div>
                                </div>
                                <div className="notify_status">
                                    <p>Action</p>
                                </div>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
            {/* End Popup */}
        </React.Fragment>
        )
    }
}
