import React, { useEffect, useState } from 'react';
import { Code, Heading1, Heading2, Heading3, Heading4, Paragraph } from "../typography/typography.jsx";

const COMPONENT_MAP = {
  'H1': Heading1,
  "H2": Heading2,
  "H3": Heading3,
  "H4": Heading4,
  "P": Paragraph,
  "PRECODE": Code, // Handle `pre` tag with `code` child specifically
};

export const RichTextComponent = ({ htmlString }) => {
  const [nodes, setNodes] = useState([]);

  useEffect(() => {
    const parser = new DOMParser();
    const doc = parser.parseFromString(htmlString, 'text/html');
    setNodes(Array.from(doc.body.childNodes));
  }, [htmlString]);

  const parseNode = (node) => {
    if (node.nodeType === 1) { // Check if it's an element node
      let tagName = node.nodeName;

      // Handle `pre` tag with `code` child specifically
      if (tagName === 'PRE' && node.firstElementChild && node.firstElementChild.nodeName === 'CODE') {
        tagName = 'PRECODE';
        const TagComponent = COMPONENT_MAP[tagName];
        if (TagComponent) {
          // Pass only the inner text of the `code` child node to the Code component
          return <TagComponent key={node.textContent}>{node.firstElementChild.textContent}</TagComponent>;
        }
      }
      else {
        const TagComponent = COMPONENT_MAP[tagName];
        if (TagComponent) { // If we have a component for this HTML tag, render it
          return <TagComponent key={node.textContent}>{node.textContent}</TagComponent>;
        }
      }
    }
    return null;
  };

  return (
    <div>
      {nodes.map((node, i) => {
        return parseNode(node);
      })}
    </div>
  );
};
