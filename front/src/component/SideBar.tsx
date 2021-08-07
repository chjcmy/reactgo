import React, {useEffect, useState} from 'react';
import {instance} from "../axios";

const SideBar = () => {

    const findUnits = async () => {
        await instance.get('/unitshosting').then(
            function (res: { data: []; }) {
                setUnit(res.data)
                console.log(res.data);
            })
            .catch(function (error: any) {
                    console.log(error)
                }
            );
    };

    const [unit, setUnit] = useState([]);

    useEffect(() => {
        findUnits();
        console.log(unit)
    }, [])
    return (
        <div>
            SideBar
        </div>
    );
};

export default SideBar;
