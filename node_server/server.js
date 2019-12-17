const { ApolloServer, gql } = require('apollo-server');
const typeDefs = gql`



type Query {
  todos: [Todo!]!
  users: [User!]!
}

type Mutation {
  createTodo(input: NewTodo!): Todo!
}


type Todo {
  id: ID!
  text: String!
  done: Boolean!
  user: User!
}

type User {
  id: ID!
  name: String!
}

input NewTodo {
  text: String!
  userId: String!
  userName: String!
}

`;

const todos = [];

const resolvers = {
  Query: {
    todos: () => todos,
  },
  Mutation:{
    createTodo: (parent, {input}) => {
      const todo = {
        text: input.text,
        done: false,
        user: {
          id: input.userId,
          name: input.userName
        }
      }
      todos.push(todo);
      return todo;
    }
  }
};


// The ApolloServer constructor requires two parameters: your schema
// definition and your set of resolvers.
const server = new ApolloServer({ typeDefs, resolvers });

// The `listen` method launches a web server.
server.listen().then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});
