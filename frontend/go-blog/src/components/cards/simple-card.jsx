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
  Text,

} from "@chakra-ui/react";
import { Link as RLink } from "react-router-dom";

// eslint-disable-next-line react/prop-types
export const SimpleCard = ({AuthorName, HeroImage, Summary, Title, WrittenOn, id}) => {
  console.log("id", id);
  return (
      <Card maxW='sm' >
        <CardBody>
          <Image
              src={HeroImage}
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
            <Link as={RLink} to={`/view/article/${id}`} onClick={()=> console.log(`/views/article/${id}`)} variant='solid' colorScheme='blue'>
          <Button>Read More</Button>
            </Link>
          </ButtonGroup>
        </CardFooter>
      </Card>
  );
};

