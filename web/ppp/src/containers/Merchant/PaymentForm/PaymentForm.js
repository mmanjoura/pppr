import React, { Component } from 'react';
import 'react-dates/initialize';
import 'react-dates/lib/css/_datepicker.css';
import { DateRangePicker } from 'react-dates';

export default class PaymentForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            acquirerId: null,
            startDate: null,
            endDate: null
        }
    }

    selectChangedHandler = (event) => {
        this.setState({
            acquirerId: event.target.value,
            startDate: this.state.startDate,
            endDate: this.state.endDate,
        }, () => {
            this.loadData(this.state.acquirerId);
        });

    }

    loadData = async () => {
        const { acquirerId, startDate, endDate } = this.state;
        console.log("Writing to paymentHandler OK!: ")
        console.log(acquirerId);
        console.log(startDate);
        console.log(endDate);
        // Get data via XHR...
        this.props.paymentHandler(acquirerId, startDate, endDate)
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
                        <label>Acquirer</label>
                        <select value={this.props.value} onChange={this.selectChangedHandler} name="acquirers" className="input-control" style={{ flex: '0.7' }}>
                            <option value="">Select an Acquirer</option>
                            <option value="d41035e2-b1bc-4bfa-a10f-de842c5b69ad">Planet-POS</option>
                            <option value="d41035e2-b1bc-4bfa-a11f-de842c5b69ad">FIPS-POS</option>
                            <option value="d41035e2-b1bc-4bfa-a12f-de842c5b69ad">OCBC-ECOM</option>
                            <option value="d41035e2-b1bc-4bfa-a13f-de842c5b69ad">FIPS-CA</option>
                            <option value="d41035e2-b1bc-4bfa-a14f-de842c5b69ad">FIPS-TRS</option>
                            <option value="d41035e2-b1bc-4bfa-a15f-de842c5b69ad">AVERY</option>
                            <option value="d41035e2-b1bc-4bfa-a16f-de842c5b69ad">GRO</option>

                        </select>
                    </div>
                    <br />
                    <br />
                    <div className="form-group">
                        <div style={{  fontSize:'10px', marginRight:'07px'}}>Toggle Approved</div> 
                        <input type="checkbox" value={this.props.value} onChange={this.CheckBoxHandler} name="acquirers" className="input-control" style={{ flex: '0.2'}} />
                        <br />
            
                    </div>

                </nav>
            </React.Fragment>
        )
    }
}



