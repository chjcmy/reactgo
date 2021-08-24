import React, { useEffect, useState } from 'react';
import styled from 'styled-components';
import { Link } from 'react-router-dom';
import {
  FormControl,
  InputLabel,
  makeStyles,
  MenuItem,
  Select,
} from '@material-ui/core';

import { CKEditor } from '@ckeditor/ckeditor5-react';
import InlineEditor from '@ckeditor/ckeditor5-editor-classic/src/classiceditor';
import Essentials from '@ckeditor/ckeditor5-essentials/src/essentials';
import Paragraph from '@ckeditor/ckeditor5-paragraph/src/paragraph';
import Bold from '@ckeditor/ckeditor5-basic-styles/src/bold';
import Italic from '@ckeditor/ckeditor5-basic-styles/src/italic';
import Underline from '@ckeditor/ckeditor5-basic-styles/src/underline';
import Strikethrough from '@ckeditor/ckeditor5-basic-styles/src/strikethrough';
import BlockQuote from '@ckeditor/ckeditor5-block-quote/src/blockquote';
import MediaEmbed from '@ckeditor/ckeditor5-media-embed/src/mediaembed';
import PasteFromOffice from '@ckeditor/ckeditor5-paste-from-office/src/pastefromoffice';
import Heading from '@ckeditor/ckeditor5-heading/src/heading';
import Font from '@ckeditor/ckeditor5-font/src/font';
import Image from '@ckeditor/ckeditor5-image/src/image';
import ImageStyle from '@ckeditor/ckeditor5-image/src/imagestyle';
import ImageToolbar from '@ckeditor/ckeditor5-image/src/imagetoolbar';
import ImageUpload from '@ckeditor/ckeditor5-image/src/imageupload';
import ImageResize from '@ckeditor/ckeditor5-image/src/imageresize';
import List from '@ckeditor/ckeditor5-list/src/list';
import Alignment from '@ckeditor/ckeditor5-alignment/src/alignment';
import Table from '@ckeditor/ckeditor5-table/src/table';
import TableToolbar from '@ckeditor/ckeditor5-table/src/tabletoolbar';
import TextTransformation from '@ckeditor/ckeditor5-typing/src/texttransformation';
import Indent from '@ckeditor/ckeditor5-indent/src/indent';
import IndentBlock from '@ckeditor/ckeditor5-indent/src/indentblock';
import Base64UploadAdapter from '@ckeditor/ckeditor5-upload/src/adapters/base64uploadadapter';
import CodeBlock from '@ckeditor/ckeditor5-code-block/src/codeblock';
import { instance } from '../../axios';

const editorConfiguration = {
  language: 'ko',
  plugins: [
    Essentials,
    Paragraph,
    Bold,
    Italic,
    Heading,
    Underline,
    Strikethrough,
    BlockQuote,
    MediaEmbed,
    PasteFromOffice,
    Font,
    Image,
    ImageStyle,
    ImageToolbar,
    ImageUpload,
    ImageResize,
    List,
    Alignment,
    Table,
    TableToolbar,
    TextTransformation,
    Indent,
    IndentBlock,
    Base64UploadAdapter,
    CodeBlock,
  ],
  toolbar: [
    'heading',
    '|',
    'bold',
    'italic',
    'underline',
    'strikethrough',
    '|',
    'fontSize',
    'fontColor',
    'fontBackgroundColor',
    '|',
    'alignment',
    'outdent',
    'indent',
    'bulletedList',
    'numberedList',
    'blockQuote',
    '|',
    'link',
    'insertTable',
    'imageUpload',
    'codeBlock',
    '|',
    'undo',
    'redo',
  ],

  alignment: {
    options: ['justify', 'left', 'center', 'right'],
  },
  table: {
    contentToolbar: ['tableColumn', 'tableRow', 'mergeTableCells'],
  },
  image: {
    resizeUnit: 'px',
    toolbar: [
      'imageStyle:alignLeft',
      'imageStyle:full',
      'imageStyle:alignRight',
      '|',
      'imageTextAlternative',
    ],
    styles: ['full', 'alignLeft', 'alignRight'],
  },
  typing: {
    transformations: {
      remove: [
        'enDash',
        'emDash',
        'oneHalf',
        'oneThird',
        'twoThirds',
        'oneForth',
        'threeQuarters',
      ],
    },
  },
  codeBlock: {
    languages: [
      { language: 'plaintext', label: 'Plain text' },
      { language: 'c', label: 'C' },
      { language: 'cs', label: 'C#' },
      { language: 'cpp', label: 'C++' },
      { language: 'css', label: 'CSS' },
      { language: 'diff', label: 'Diff' },
      { language: 'html', label: 'HTML' },
      { language: 'java', label: 'Java' },
      { language: 'javascript', label: 'JavaScript' },
      { language: 'php', label: 'PHP' },
      { language: 'python', label: 'Python' },
      { language: 'ruby', label: 'Ruby' },
      { language: 'typescript', label: 'TypeScript' },
      { language: 'xml', label: 'XML' },
      { language: 'go', label: 'Go' },
    ],
  },
  placeholder: '글을 입력해보세요!',
};

const Outer = styled.div`
  margin: 2% 1% 0 2%;
  width: 80%;
`;

const FamilyFont = styled.div`
  font-family: Neodgm, serif;
`;

const Creator = styled.div`
  display: flex;
  height: 50%;
`;
const Save = styled(Link)`
  display: flex;
  height: 50%;
`;

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
const BookWrite = () => {
  const classes = useStyles();
  const [units, setUnits] = useState([]);
  const [read, setRead] = useState(false);
  const [chapter, setChapter] = useState('');
  const [title, setTitle] = useState('');
  const [subject, setSubject] = useState('');
  const [open, setOpen] = React.useState(false);
  const findUnits = async () => {
    await instance
      .get('/unitshosting')
      .then((res) => {
        setUnits(res.data);
      })
      .catch(() => {});
  };

  const createBook = async () => {
    await instance.post(
      '/bookcreate',
      {
        unit: chapter,
        title,
        subject,
      },
      {
        headers: {
          'Content-Type': 'application/json',
          id: sessionStorage.getItem('id'),
        },
      }
    );
  };

  if (read) {
    setTimeout(() => setRead(false), 500);
  }
  useEffect(() => {
    findUnits().then();
  }, []);
  const click = () => {
    createBook().then();
  };

  const handleClose = () => {
    setOpen(false);
  };

  const handleOpen = () => {
    setOpen(true);
  };

  const handleChange = (event) => {
    setChapter(event.target.value);
  };

  return (
    <>
      <Outer>
        <FamilyFont>
          <Creator>
            <div>
              <FormControl
                style={{ width: '100%' }}
                className={classes.formControl}
              >
                <InputLabel id="demo-controlled-open-select-label">
                  unit
                </InputLabel>
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
                  {units.map((unit) => (
                    <MenuItem key={unit.id} value={unit.id}>
                      {unit.content_name}
                    </MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>
            <Save
              to="/"
              onClick={click}
              style={{ marginLeft: '30%', width: '100%' }}
            >
              <button
                type="button"
                className="nes-btn is-primary"
                style={{ marginLeft: '5%', fontSize: 'xx-large' }}
              >
                저장
              </button>
            </Save>
          </Creator>
          <div className="nes-field" style={{ marginBottom: '2%' }}>
            {/* eslint-disable-next-line jsx-a11y/label-has-associated-control */}
            <label htmlFor="name_field">주제</label>
            <input
              type="text"
              id="name_field"
              className="nes-input"
              style={{ width: '90%', height: '50%' }}
              onChange={(e) => setTitle(e.target.value)}
            />
          </div>
        </FamilyFont>
        <CKEditor
          editor={InlineEditor}
          data="<p>Hello from CKEditor 5!</p>"
          config={editorConfiguration}
          onChange={(event, editor) => {
            const data = editor.getData();
            setSubject(data);
            setRead(true);
          }}
        />
      </Outer>
    </>
  );
};
export default BookWrite;
