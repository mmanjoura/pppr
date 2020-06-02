import React, { Component } from 'react'
import 'react-dates/initialize'
import 'react-dates/lib/css/_datepicker.css'
import { DateRangePicker } from 'react-dates';

export default class ChartForm extends Component {
    constructor(props) {
        super(props)
        this.state = {
          chartType: 1
        }
    }
    selectChangedHandler = (event) => {

        this.setState({ chartType: event.target.value });
        this.loadData(event.target.value);
    }
    loadData = async (currentValue) => {
        // Get data via XHR...
        this.props.chartChangedHandler(currentValue)
        console.log("Child Called Handler Successfuly", currentValue )
    }
    render() {
        return (
            <React.Fragment>
                <nav className="panel-left">
                    <div className="form-group">
                        <label></label>
                        <DateRangePicker
                            startDate={this.state.startDate} // momentPropTypes.momentObj or null,
                            startDateId="your_unique_start_date_id" // PropTypes.string.isRequired,
                            endDate={this.state.endDate} // momentPropTypes.momentObj or null,
                            endDateId="your_unique_end_date_id" // PropTypes.string.isRequired,
                            onDatesChange={({ startDate, endDate }) => this.setState({ startDate, endDate })} // PropTypes.func.isRequired,
                            focusedInput={this.state.focusedInput} // PropTypes.oneOf([START_DATE, END_DATE]) or null,
                            onFocusChange={focusedInput => this.setState({ focusedInput })} // PropTypes.func.isRequired,
                        />
                    </div>
                    <br />
                    <div className="form-group">
                        <label>Chart Type</label>
                        <select value={this.props.value} onChange={this.selectChangedHandler}  name="charts" className="input-control" style={{ flex: '0.7' }}>
                            <option value="1">Amount-Fees</option>
                            <option value="2">Rebate-Acquirer</option>
                            <option value="3">Acquirer-Transactions</option>
                            <option value="4">DCC-Local</option>
                            <option value="5">Transactions-Region</option>
                            <option value="6">ChargeBack-Region</option>
                   
                        </select>
                    </div>                    
                </nav>
            </React.Fragment>
        )
    }
}


