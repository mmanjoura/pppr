import React, { Component } from 'react';
import 'react-dates/initialize';
import 'react-dates/lib/css/_datepicker.css';
import { DateRangePicker } from 'react-dates';


export default class RateForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            startDate: null,
            endDate: null
        }
    }
    alertStartDate = () => {
        alert(this.state.startDate)
    }
    alertEndtDate = () => {
        alert(this.state.endDate)
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
                        <label>Rates</label>
                        <select id="cars" name="rate" className="input-control" style={{ flex: '0.7' }}>
                            <option>All</option>
                            <option>USD</option>
                            <option>GBP</option>
                            <option>SGD</option>
                            <option>EUR</option>
                            <option>CNY</option>
                   
                        </select>
                    </div>                    

                    {/* <button onClick={this.alertStartDate}>start date</button>
                    <button onClick={this.alertEndtDate}>end date</button> */}
                </nav>
            </React.Fragment>
        )
    }
}



