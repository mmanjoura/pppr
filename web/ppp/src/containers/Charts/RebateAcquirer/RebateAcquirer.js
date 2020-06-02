import React, { PureComponent } from 'react';
import {
    LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ReferenceLine,
} from 'recharts';
import Controls from '../../../components/Content/Controls/Controls'

const data = [
    {
        name: 'Planet-POS', Rebate: 4000, TrxsAmount: 2400, amt: 2400,
    },
    {
        name: 'FIPS-POS', Rebate: 3000, TrxsAmount: 1398, amt: 2210,
    },
    {
        name: 'OCBC-ECOM', Rebate: 2000, TrxsAmount: 9800, amt: 2290,
    },
    {
        name: 'FIPS-CA', Rebate: 2780, TrxsAmount: 3908, amt: 2000,
    },
    {
        name: 'FIPS-TRS', Rebate: 1890, TrxsAmount: 4800, amt: 2181,
    },
    {
        name: 'AVERY', Rebate: 2390, TrxsAmount: 3800, amt: 2500,
    },
    {
        name: 'GRO', Rebate: 3490, TrxsAmount: 4300, amt: 2100,
    },
];

export default class RebateAcquirer extends PureComponent {

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
                <LineChart
                    width={800}
                    height={500}
                    data={data}
                    margin={{
                        top: 20, right: 50, left: 20, bottom: 5,
                    }}
                >
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="name" />
                    <YAxis />
                    <Tooltip />
                    <Legend />
                    <ReferenceLine x="OCBC-ECOM" stroke="red" label="Max Rebate" />
                    <ReferenceLine y={9800} label="Max" stroke="red" />
                    <Line type="monotone" dataKey="TrxsAmount" stroke="#8884d8" />
                    <Line type="monotone" dataKey="Rebate" stroke="#82ca9d" />
                </LineChart>
            </React.Fragment>

        );
    }
}
