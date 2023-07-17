import SyntaxHighlighter from 'react-syntax-highlighter';
import {atomOneDark, nord} from 'react-syntax-highlighter/dist/esm/styles/hljs';
import {Box, Heading, Text} from "@chakra-ui/react";


export const Heading1 = ({children}) => {
  return <Heading as='h1' fontSize='3xl' mb={2} align='center'>{children}</Heading>;
};

export const Heading2 = ({children}) => {
  return <Text as='h2' fontSize='2xl' mb={2} align='center'>{children} </Text>;
};

export const Heading3 = ({children}) => {
  return <Text fontSize='xl'>{children}</Text>;
};

export const Heading4 = ({children}) => {
  return <Text fontSize='md' align='left'>{children}</Text>;
};


export const Paragraph = ({children}) => {
  return <Text fontSize='sm' align='left'>{children}</Text>;
};


export const Code = ({children}) => {
  console.log("htmlString", children);

  return (
      <Box mt={2} mb={2}>
        <SyntaxHighlighter language="javascript" style={nord}>
          {children}
        </SyntaxHighlighter>
      </Box>
  );
};

