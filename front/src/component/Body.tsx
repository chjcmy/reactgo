import React from 'react';
import styles from '../CSS/header.css';
import classNames from "classnames/bind";
const css = classNames.bind(styles);

const Body = () => {
    return (
        <div className={css('Body')}>
            Body
        </div>
    );
};

export default Body;
