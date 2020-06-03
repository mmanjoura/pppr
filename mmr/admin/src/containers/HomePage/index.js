import React, { memo } from 'react';
import { get, upperFirst } from 'lodash';
import { auth } from 'strapi-helper-plugin';

import { Block, Container } from './components';

const HomePage = ({ global: { plugins }, history: { push } }) => {
    const username = get(auth.getUserInfo(), 'username', '');
    return (
        <>
            <Container className="container-fluid">
                <div className="row">
                    <div className="col-12">
                        <Block>Hello {upperFirst(username)} what can I do for you today!</Block>
                    </div>
                </div>
            </Container>
        </>
    );
};

export default memo(HomePage);