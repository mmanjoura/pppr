import React, { Component } from 'react';
import axios from 'axios';
import Log from '../../components/Content/Log/Log';
import Header from '../../components/Content/Log/Header/Header';
import LogForm from './LogForm/LogForm';
import { ModalSpinner, Modal } from '../../components/Layout/Modal/Modal';



export default class LogMessages extends Component {

    constructor(props) {
        super(props)
        this.logMessagesHandler = this.logMessagesHandler.bind(this)
        this.state = {
            logs: [],
            filteredLogs: [],
            showSummary: false,
            level: '',
            loading: false,
            showHeader: true
        }
    }

    componentDidMount() {
        var todayDate = new Date().toISOString().slice(0, 10);
        axios.get('http://localhost:8040/logs/' + todayDate)
            .then(response => {
                this.setState({
                    logs: response.data,
                });
                console.log(this.state.logs)
            });
    }

    showSummaryOnHoverHandler = () => {
        this.setState({
            showSummary: false,
        });
    }

    closeModalHandler = () => {
        this.setState({ showSummary: false });
    }

    logMessagesHandler(level) {
        this.setState({
            viewPaymentDetails: true,
            level: level,
            loading: true,
        }, () => {
            const result = this.state.logs.filter(word => word.level == level)
            this.setState({ filteredLogs: result, loading: true }, () => {
                this.setState({ loading: false })
            });

        });
    }

    render = () => {
        let logs = this.state.logs.map((p, index) => {

            return <Log
                key={index}
                createddate={p.createddate}
                createdtime={p.createdtime}
                level={p.level}
                servicename={p.servicename}
                callingmethod={p.callingmethod}
                host={p.host}
                latency={p.latency}
                message={p.body}
                showSummaryOnHover={() => this.showSummaryOnHoverHandler()}
                show={this.state.showSummary}

            />
        });


        if (this.state.loading) {
            logs = <ModalSpinner />
        } else {
            logs = this.state.filteredLogs.map((p, index) => {

                return <Log
                    key={index}
                    createddate={p.createddate}
                    createdtime={p.createdtime}
                    level={p.level}
                    servicename={p.servicename}
                    callingmethod={p.callingmethod}
                    host={p.host}
                    latency={p.latency}
                    message={p.body}
                    showSummaryOnHover={() => this.showSummaryOnHoverHandler()}
                    show={this.state.showSummary}

                />
            });
        }

        return (
            <React.Fragment>
                <LogForm logMessagesHandler={this.logMessagesHandler} />
                <br></br>
                {this.state.showHeader == true ? <Header /> : null}
                <div id="scroller">
                    {logs}
                </div>

            </React.Fragment>

        )
    }
}
