import React, { PureComponent } from 'react';
import {
  BarChart, Bar, Brush, ReferenceLine, XAxis, YAxis, CartesianGrid, Tooltip, Legend,
} from 'recharts';
import Controls from '../../../components/Content/Controls/Controls'

const data = [
  { name: '1', Amount: 300, Fees: 456 },
  { name: '2', Amount: -145, Fees: 230 },
  { name: '3', Amount: -100, Fees: 345 },
  { name: '4', Amount: -8, Fees: 450 },
  { name: '5', Amount: 100, Fees: 321 },
  { name: '6', Amount: 9, Fees: 235 },
  { name: '7', Amount: 53, Fees: 267 },
  { name: '8', Amount: 252, Fees: -378 },
  { name: '9', Amount: 79, Fees: -210 },
  { name: '10', Amount: 294, Fees: -23 },
  { name: '12', Amount: 43, Fees: 45 },
  { name: '13', Amount: -74, Fees: 90 },
  { name: '14', Amount: -71, Fees: 130 },
  { name: '15', Amount: -117, Fees: 11 },
  { name: '16', Amount: -186, Fees: 107 },
  { name: '17', Amount: -16, Fees: 926 },
  { name: '18', Amount: -125, Fees: 653 },
  { name: '19', Amount: 222, Fees: 366 },
  { name: '20', Amount: 372, Fees: 486 },
  { name: '21', Amount: 182, Fees: 512 },
  { name: '22', Amount: 164, Fees: 302 },
  { name: '23', Amount: 316, Fees: 425 },
  { name: '24', Amount: 131, Fees: 467 },
  { name: '25', Amount: 291, Fees: -190 },
  { name: '26', Amount: -47, Fees: 194 },
  { name: '27', Amount: -415, Fees: 371 },
  { name: '28', Amount: -182, Fees: 376 },
  { name: '29', Amount: -93, Fees: 295 },
  { name: '30', Amount: -99, Fees: 322 },
  { name: '31', Amount: -52, Fees: 246 },
  { name: '32', Amount: 154, Fees: 33 },
  { name: '33', Amount: 205, Fees: 354 },
  { name: '34', Amount: 70, Fees: 258 },
  { name: '35', Amount: -25, Fees: 359 },
  { name: '36', Amount: -59, Fees: 192 },
  { name: '37', Amount: -63, Fees: 464 },
  { name: '38', Amount: -91, Fees: -2 },
  { name: '39', Amount: -66, Fees: 154 },
  { name: '40', Amount: -50, Fees: 186 },
];

export default class AmountFee extends PureComponent {

  state = {
        logs: [],
        showSummary: false,
        formType: 5,
    }

  showSummaryOnHoverHandler = () => {
    this.setState({
      showSummary: false,
    });
  }

  closeModalHandler = () => {
    this.setState({ showSummary: false });
  }

  render() {
    return (
      <React.Fragment>
        <BarChart
          width={800}
          height={500}
          data={data}
          margin={{
            top: 5, right: 30, left: 20, bottom: 5,
          }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="name" />
          <YAxis />
          <Tooltip />
          <Legend verticalAlign="top" wrapperStyle={{ lineHeight: '40px' }} />
          <ReferenceLine y={0} stroke="#000" />
          <Brush dataKey="name" height={30} stroke="#8884d8" />
          <Bar dataKey="Amount" fill="#8884d8" />
          <Bar dataKey="Fees" fill="#82ca9d" />
        </BarChart>
      </React.Fragment>

    );
  }
}
