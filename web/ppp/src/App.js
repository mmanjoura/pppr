import React from 'react';
import axios from 'axios';
import Layout from './components/Layout/Layout';
import Right from './components/Layout/Navigation/Right/Right'
import Left from './components/Layout/Navigation/Left/Left'
import Popup from './components/Layout/Navigation/Popup/Popup'
import Notification from './components/Layout/Navigation/Notification/Notification'
import Tools from './components/Layout/Navigation/Tools/Tools'

import { BrowserRouter } from 'react-router-dom'
import PlanetLogo from './assets/images/planet-logo.png'
import Content from './containers/Content/Content'
;


import $ from 'jquery';


class App extends React.Component {

  componentDidMount() {  

    $(document).ready(function () {
      $(".profile .icon_wrap").click(function () {
        $(this).parent().toggleClass("active");
        $(".notifications").removeClass("active");
      });

      $(".notifications .icon_wrap").click(function () {
        $(this).parent().toggleClass("active");
        $(".profile").removeClass("active");
      });

      $(".show_all .link").click(function () {
        $(".notifications").removeClass("active");
        $(".popup").show();
      });

      $(".close").click(function () {
        $(".popup").hide();
      });
    });

  }

  render() {
    return (
      <BrowserRouter>
        <Layout>
          <div className="wrapper">
            <div className="navbar">
              <Left logo={PlanetLogo} />
              <div className="navbar_right">
                <Right logo={PlanetLogo} />
                   <Notification />
              
                <Tools />
              </div>

            </div>
            <div className="content-container">
              <article className="main-content" style={{ flex: '.885' }}>
                  <Popup />                
                <Content />
              </article>
            </div>
          </div>
        </Layout>
      </BrowserRouter>

    );

  }
}
export default App;
