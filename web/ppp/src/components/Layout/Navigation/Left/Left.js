import React from 'react'
import { Link } from 'react-router-dom'

const Left = (props) => {
    return (
        <React.Fragment>
            <div className="navbar_left">
            <Link className="address" to="/charts"> <img src={props.logo} width="55px" height="55px"/></Link>
            </div>
        </React.Fragment>
    )
}
export default Left
