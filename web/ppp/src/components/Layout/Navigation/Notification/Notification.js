import React, { Component } from 'react'
import axios from 'axios';
import NotificationSystem from 'react-notification-system';
import { withRouter } from "react-router-dom";
import { Link } from 'react-router-dom'


class Notification extends Component {

    constructor(props) {
        super(props)
        this.state = {
            notifications: [],
            clicks: 0,
        }
    }

    async componentDidMount() {

        try {
            setInterval(async () => {

               await  axios.get('http://localhost:8040/states/false')
                    .then(response => {
                        this.setState({
                            notifications: response.data,
                        });
                        console.log(this.state.notifications)
                    });

            }, 10000);
        } catch (e) {
            console.log(e);
        }

    }

    notificationSystem = React.createRef();
    addNotification = event => {
        this.setState({ clicks: this.state.clicks + 1 })
        console.log(this.state.clicks)
        event.preventDefault();
        const notification = this.notificationSystem.current;

        if (this.state.notifications.length == 0) {
            notification.addNotification({
                title: "No Action Needed",
                message: "All Payments have been approved",
                level: 'success',
                position: 'br',
            });
        }
        else {
            notification.addNotification({
                title: "Payment Ready for Approval",
                message: "",
                level: 'warning',
                position: 'tc',
                children: (
                    <div>
                        <li><Link className="address" to="/payments"><span className="picon"><i className="" /></span>FIPS-CA</Link></li>
                        <li><Link className="address" to="/payments"><span className="picon"><i className="" /></span>2020-05-2020</Link></li>
                    </div>
                )
            });
        }
    };

    render() {
        return (
            <React.Fragment>
                <div className="notifications">
                    <div className="icon_wrap"><i className="far fa-bell" onClick={this.addNotification} />
                        {this.state.notifications.length > 0 ? <div className="icon-badge">{this.state.notifications.length}</div> : null}
                        <div className="icon-badge">{this.state.notifications.length}</div>
                        <NotificationSystem ref={this.notificationSystem} />
                    </div>
                </div>

            </React.Fragment>
        )
    }
}

export default withRouter(Notification)
