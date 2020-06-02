import React, { Component } from 'react'
import axios from 'axios'
import Report from '../../components/Content/Report/Report';
import Header from '../../components/Content/Report/Header/Header'
import Csv from '../../assets/images/csv.png'
import Pdf from '../../assets/images/pdf.png'
import Xls from '../../assets/images/xls.png'
import Approved from '../../assets/images/approved.png'
import ReportForm from './ReportForm/ReportForm';
import { ModalSpinner, Modal } from '../../components/Layout/Modal/Modal';



export default class MerchantPayments extends Component {

    constructor(props) {
        super(props)
        this.reportHandler = this.reportHandler.bind(this)

        this.state = {
            reports: [],
            filteredReports: [],
            showSummary: false,
            loading: false,
            acquirerId: '',
            startDate: '',
            endDate: '',
            showHeader: false,
        }
    }

    // Set the AcquirerId and Date in payment CMP
    reportHandler(acquirerId, startDate, endDate) {
        this.setState({
            acquirerId: acquirerId,
            startDate: startDate,
            endDate: endDate,
            loading: true,
        }, () => {

            axios.get('http://localhost:8040/acquirer/reports/' + acquirerId)
                .then(response => {
                    this.setState({
                        reports: response.data,
                    });
                    console.log("Select AcquirerId is: ", acquirerId)
                    // Acquirer Drop down changed
                    const result = this.state.reports.filter(word => word.acquirerid == acquirerId)
                    this.setState({ filteredReports: result, loading: true }, () => {
                        this.setState({ loading: false })
                    });

                });
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
    render() {

        const reports = this.state.reports.map((p, index) => {

            return <Report
                key={index}
                merchantname={p.merchantname}
                merchantid={p.merchantid}
                paymentid={p.paymentid}
                pdffile={p.filepath}
                xlsfile={p.filepath}
                csvfile={p.filepath}
                csv={Csv}
                pdf={Pdf}
                xls={Xls}
                approved={Approved}
                showSummaryOnHover={() => this.showSummaryOnHoverHandler()}
                show={this.state.showSummary}
            />
        });

       

        if (this.state.loading) {
            return (
                <React.Fragment>
                    <ReportForm reportHandler={this.reportHandler} />
                    <ModalSpinner />
                </React.Fragment>
            )
        } else {

            return (
                <React.Fragment>
                    <ReportForm reportHandler={this.reportHandler} />
                    <Header />
                    <div id="scroller">
                        {reports}
                    </div>
                </React.Fragment>
            )
        }
    }
}


