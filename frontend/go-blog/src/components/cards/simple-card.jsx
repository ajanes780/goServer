import {
  Button,
  ButtonGroup,
  Card,
  CardBody,
  CardFooter,
  Divider,
  Heading,
  Image,
  Link,
  Stack,
  Text
} from "@chakra-ui/react";

// eslint-disable-next-line react/prop-types
export const SimpleCard = ({AuthorName, HeroImage, Summary, Title, WrittenOn, key}) => {
  return (
      <Card maxW='sm' key={key}>
        <CardBody>
          <Image
              src="./src/assets/chakra.png"
              alt='Green double couch with wooden legs'
              borderRadius='lg'
          />
          <Stack mt='6' spacing='3'>
            <Heading size='md'>{Title}</Heading>
            <Text textAlign='left'>
              {Summary}
            </Text>
          </Stack>
          <Stack mt='6' spacing='3'>
            <Text fontSize='sm' color='gray.500'>
              By:{AuthorName}

            </Text>
            <Text fontSize='sm' color='gray.500'>
              Published: {WrittenOn}
            </Text>
          </Stack>
        </CardBody>
        <Divider/>
        <CardFooter>
          <ButtonGroup spacing='2'>
            <Link as={Button} variant='solid' colorScheme='blue'>
              Read more
            </Link>
          </ButtonGroup>
        </CardFooter>
      </Card>
  );
};

