import React, { Component } from 'react';
import 'react-dates/initialize';
import 'react-dates/lib/css/_datepicker.css';
import { DateRangePicker } from 'react-dates';


export default class LogForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            startDate: null,
            endDate: null,
            level: null,
        }
    }

    selectChangedHandler = (event) => {

        this.setState({ level: event.target.value });
        this.loadData(event.target.value);
    }
    loadData = async (currentValue) => {
        // Get data via XHR...
        this.props.logMessagesHandler(currentValue)
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
                        <label>Level</label>
                        <select value={this.props.value} onChange={this.selectChangedHandler} name="log" className="input-control" style={{ flex: '0.7' }}>
                            <option value="ERROR">ERROR</option>
                            <option value="INFO">INFO</option>
                            <option value="WARNING">WARNING</option>
                            
                        </select>
                    </div>
                </nav>
            </React.Fragment>
        )
    }
}



