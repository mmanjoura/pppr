import React from 'react'
import Backdrop from '../Backdrop/Backdrop'


const Modal = (props) => {
    console.log(props.show)
    return (        
            <React.Fragment>
                <Backdrop show={ props.show } clicked={props.modalClosed} />
                <div className="Modal" style={{transform: props.show ? 'translateY(0)' : 'translateY(-100vh)'}}>
                    { props.children }
                </div>
            </React.Fragment>
            )
}



export default Modal
const ModalSpinner = () => {
    return (
        <React.Fragment>
            <div className="Loader">Loading...</div>
        </React.Fragment>
    )
}

export {
    Modal,
    ModalSpinner,
}