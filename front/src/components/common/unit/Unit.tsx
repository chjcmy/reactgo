import React, {useEffect, useState} from 'react';
import {instance} from "../../../axios";

const Unit = ({match} : {match: any})=> {

    const [books, setBooks] = useState<any[]>([]);

    const findAllBook = async () => {
        const res = await instance.get('/bookshow/0')
        setBooks(res.data)
        console.log(res.data)
    };

    const findUnitBook = async () => {
        const res = await instance.get(`/pickunitbooks/${match.params.unit}/1`)
        setBooks(res.data)
        console.log(books)
    };

    useEffect(() => {
       if (!match.params.unit) {
           findAllBook()
       } else {
           findUnitBook()
       }
    }, [match.params.unit]);

    return (
        <div>
            {match.params.unit}
            {books.map(books => (
                    <div key={books.id}>
                        {books.subject}
                    </div>
                )
            )}
        </div>
    );
};

export default Unit;
