import React, { useEffect, useState } from 'react';
import 'nes.css/css/nes.min.css';
// eslint-disable-next-line import/order
import { instance } from '../../axios';
import styled from 'styled-components';
import './ckcontent.css';
import { CKEditor } from '@ckeditor/ckeditor5-react';

import ClassicEditor from '@ckeditor/ckeditor5-editor-classic/src/classiceditor';
import Bold from '@ckeditor/ckeditor5-basic-styles/src/bold';
import Italic from '@ckeditor/ckeditor5-basic-styles/src/italic';
import Essentials from '@ckeditor/ckeditor5-essentials/src/essentials';
import Paragraph from '@ckeditor/ckeditor5-paragraph/src/paragraph';
import Heading from '@ckeditor/ckeditor5-heading/src/heading';
import Underline from '@ckeditor/ckeditor5-basic-styles/src/underline';
import Strikethrough from '@ckeditor/ckeditor5-basic-styles/src/strikethrough';
import BlockQuote from '@ckeditor/ckeditor5-block-quote/src/blockquote';
import MediaEmbed from '@ckeditor/ckeditor5-media-embed/src/mediaembed';
import PasteFromOffice from '@ckeditor/ckeditor5-paste-from-office/src/pastefromoffice';
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
import GoogleAdsense from 'react-adsense-google';

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
};

const BookDiv = styled.div`
  margin: 2% 2% 2% 2%;
  width: 100%;
  display: block;
`;

const Date = styled.div`
  font-size: large;
`;

const Book = ({ match }) => {
  const [rbook, setRbook] = useState([]);

  useEffect(() => {
    const findBook = async () => {
      const res = await instance.get(`/bookread/${match.params.id}`);
      setRbook(res.data);
    };

    findBook().then();
  }, [match.params.id, match.params.unit]);

  return (
    <>
      <BookDiv>
        <GoogleAdsense adSlot="9876543210" adClient="ca-pub-7458640452724959" />
        <div style={{ fontFamily: 'Neodgm', fontSize: 'xx-large' }}>
          ??????:
          {rbook.title}
        </div>
        <div style={{ fontFamily: 'Neodgm' }}>
          <div style={{ fontSize: 'x-large' }}>
            ?????????:
            {rbook.edges?.userid.name}
          </div>
        </div>
        <div style={{ fontSize: 'x-large' }}>
          ?????????: {rbook.edges?.unitid.content_name}
          <Date>
            ?????? ??????:
            {rbook.create_at}
          </Date>
          <Date>
            ??????????????? ??????:
            {rbook.updated_at}
          </Date>
        </div>
        <CKEditor
          editor={ClassicEditor}
          disabled
          config={editorConfiguration}
          data={rbook.subject}
        />
      </BookDiv>
    </>
  );
};
export default Book;
