import React, { Component } from 'react';
import 'react-dates/initialize';
import 'react-dates/lib/css/_datepicker.css';
import { DateRangePicker } from 'react-dates';


export default class RateForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            currency: '',
            startDate: null,
            endDate: null,          
        }
    }

    selectChangedHandler = (event) => {
        this.setState({
            currency: event.target.value,
            startDate: this.state.startDate,
            endDate: this.state.endDate,
        }, () => {
            this.loadData(this.state.currency);
        });
    }

    loadData = async () => {
        const { currency, startDate, endDate } = this.state;
        // Get data via XHR...
        this.props.ratesHandler(currency, startDate, endDate)
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
                        <select value={this.props.value} onChange={this.selectChangedHandler} name="rate" className="input-control" style={{ flex: '0.7' }}>
                            <option value="">Select Rate</option>
                            <option value="All">All</option>
                            <option value="USD">USD</option>
                            <option value="GBP">GBP</option>
                            <option value="SGD">SGD</option>
                            <option value="EUR">EUR</option>
                            <option value="CNY">CNY</option>

                        </select>
                    </div>

                </nav>
            </React.Fragment>
        )
    }
}



