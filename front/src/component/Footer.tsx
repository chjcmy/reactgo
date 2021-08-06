import React from 'react';
import classNames from "classnames/bind";
import styles from '../CSS/header.css';
const css = classNames.bind(styles);
const Footer = () => {
    return (
        <div className={css('FooterBody')}>
            Footer
        </div>
    );
};

export default Footer;
