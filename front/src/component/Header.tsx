import React from 'react';
import styles from '../CSS/header.css';
import classNames from "classnames/bind";
const css = classNames.bind(styles);

const Header = () => {

    return (
        <div className={css('headerBody')}>
            <h1>Sung.Blog</h1>
        </div>
    );

};

export default Header;
