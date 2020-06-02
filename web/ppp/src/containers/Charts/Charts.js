import React, { Component } from 'react'
import RebateAcquirer from './RebateAcquirer/RebateAcquirer'
import AmountFee from './AmountFee/AmountFee'
import ChartForm from './ChartForm/ChartForm'

export default class Charts extends Component {

    constructor(props) {
        super(props)
        this.chartChangedHandler = this.chartChangedHandler.bind(this)
        this.state = {
            chartType: 1

        }
    }

    chartChangedHandler(type) {
        this.setState({
            chartType: type,
        }, () => {
            console.log("We came back here!")
        });
    }

    render() {
        switch (this.state.chartType) {
            case 1:
                return (
                    <React.Fragment>
                        <ChartForm chartChangedHandler={this.chartChangedHandler} />
                        <AmountFee />
                    </React.Fragment>
                );

            case 2:
                return (
                    <React.Fragment>
                        <ChartForm chartChangedHandler={this.chartChangedHandler} />
                        <RebateAcquirer />
                    </React.Fragment>
                );
            default:
                return (
                    <React.Fragment>
                        <ChartForm chartChangedHandler={this.chartChangedHandler} />
                        <RebateAcquirer />
                    </React.Fragment>
                );
        }

    }
}
