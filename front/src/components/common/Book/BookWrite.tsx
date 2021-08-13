import React, {FC, useEffect, useState} from 'react';
import {instance} from "../../../axios";
import styled from "styled-components";
import SunEditor from "suneditor-react";
import 'suneditor/dist/css/suneditor.min.css';
import plugins from 'suneditor/src/plugins'
import {Link} from "react-router-dom";
import axios from "axios";
import {FormControl, InputLabel, makeStyles, MenuItem, Select} from "@material-ui/core";

const Subject = styled.div`
  margin-top: 2%;
  font-family: Neodgm, serif;
  color: #000000;
`;

const Outer = styled.div`
  position: inherit;
  margin: 2% 1% 0 2%;
  width: 80%;
  font-family: Neodgm, serif;
`

const Creator = styled.div`
  display: flex;
  height: 50%;
`

const Save = styled(Link)`
  display: flex;
  height: 50%;
`

const useStyles = makeStyles((theme) => ({
    button: {
        display: 'block',
        marginTop: theme.spacing(2),
    },
    formControl: {
        margin: theme.spacing(1),
        minWidth: 120,
    },
}));


const BookWrite: FC = () => {

    const classes = useStyles();

    const [units, setUnits] = useState<any[]>([]);
    const [read, setRead] = useState<boolean>(false);
    const [chapter, setChapter] = useState('');
    const [title, setTitle] = useState('');
    const [subject, setSubject] = useState('');
    const [open, setOpen] = React.useState(false);


    const findUnits = async () => {
        await instance.get('/unitshosting').then(
            function (res: { data: []; }) {
                setUnits(res.data)
                console.log(units)
            })
            .catch(function (error: any) {
                    console.log(error)
                }
            );
    };

    const createBook = async () => {
        console.log(chapter,title,subject);
        await axios.post('http://localhost:8000/bookcreate', {
            unit: chapter,
            title: title,
            subject: subject,
            }, {
            headers: {
                'Content-Type': 'application/json',
                'id': localStorage.getItem("id")
            }
        })
    };

    if (read) {
        setTimeout(()=> setRead(false),500)
    }

    useEffect(() => {
        findUnits().then()
    }, []);

    const click = () => {
        console.log(chapter,title,subject);
        createBook().then();
    }

    const handleClose = () => {
        setOpen(false);
    };

    const handleOpen = () => {
        setOpen(true);
    };

    const handleChange = (event: any) => {
        setChapter(event.target.value);
    };

    return <>
        <Outer>
            <Creator>
                <div>
                <FormControl style={{width:"100%"}} className={classes.formControl}>
                    <InputLabel id="demo-controlled-open-select-label">unit</InputLabel>
                    <Select
                        labelId="demo-controlled-open-select-label"
                        id="demo-controlled-open-select"
                        open={open}
                        onClose={handleClose}
                        onOpen={handleOpen}
                        value={chapter}
                        onChange={handleChange}
                    >
                        <MenuItem value="">
                            <em>None</em>
                        </MenuItem>
                        {units.map(unit => <MenuItem value={unit.id}>{unit.content_name}</MenuItem>
                        )}
                    </Select>
                </FormControl>
                </div>
                { read ?
                    <span className="nes-text is-success" style={{fontSize:"xxx-large", marginRight:"2%", marginTop:"2%", marginLeft:"3%"}}>읽혀짐</span>
                    :
                    null
                }
               <Save to="/" onClick={click} style={{marginLeft:"30%", width:"100%"}}>
                    <button type="button" className="nes-btn is-primary" style={{marginLeft:"2%", fontSize:"xx-large"}}>저장</button>
                </Save>
            </Creator>
            <div className="nes-field" style={{marginBottom: "2%"}}>
                <label htmlFor="name_field">주제</label>
                <input type="text" id="name_field" className="nes-input" style={{width: "90%", height: "50%"}}
                       onChange={e => setTitle(e.target.value)}/>
            </div>
            {/* eslint-disable-next-line @typescript-eslint/no-unused-expressions */}
            <SunEditor onChange={content => {setSubject(content), setRead(true), click}} setOptions={{
                plugins: plugins,
                buttonList: [['undo', 'redo'],
                    ['font', 'fontSize', 'formatBlock'],
                    ['paragraphStyle', 'blockquote'],
                    ['bold', 'underline', 'italic', 'strike', 'subscript', 'superscript'],
                    ['fontColor', 'hiliteColor', 'textStyle'],
                    ['removeFormat'],
                    '/',
                    ['outdent', 'indent'],
                    ['align', 'horizontalRule', 'list', 'lineHeight'],
                    ['table', 'link', 'image', 'video', 'audio'],
                    ['fullScreen', 'showBlocks', 'codeView'],
                    ['preview', 'print'],
                    ['save', 'template']
                ]
            }}/>
        </Outer>
    </>;
};

export default BookWrite;
