# Software Architecture — Inversion of Control Part 2
## Unleash the Power of Decoupled Design: The Essential Guide to the Service Locator Pattern for Software Developers
![hero image ](/images/inversion-of-control-part2.png)

The Service Locator pattern is a design pattern used in software development to decouple the creation of objects from their use. Ideally, you will use this pattern in applications relying on dependency injection, where things depend on other objects to perform their intended functions. The Service Locator pattern is part of the Inversion of Control (IoC) design pattern, which helps to manage these dependencies in a clean and organized way.

The Service Locator pattern uses a centralized location to store and retrieve objects used by an application. This location is known as the “service locator,” It acts as an intermediary between the client object that needs the service and the actual service object that provides the functionality.

When a client object needs a service, it asks the service locator for the required service. The service locator then either creates a new instance of the service or retrieves a previously created instance from its cache and returns it to the client object.

This pattern has several advantages. First, it helps reduce the coupling between client objects and service objects, as the client objects do not need to know the details of how the services are created or implemented. The service locator pattern makes it easier to change the implementation of a service without affecting the client objects that depend on it.

Second, the Service Locator pattern can improve the performance of an application by reusing previously created instances of services instead of creating new instances every time they are needed.

The Service Locator pattern can be implemented in several different ways, depending on the application’s specific requirements. One typical implementation is to use a singleton class to serve as the service locator. This class stores a map of service names to service objects and provides methods for registering and retrieving services.

```js

class Service {
  execute() {
    throw new Error("This method must be implemented");
  }
}

class Service1 extends Service {
  execute() {
    console.log("Executing Service1");
  }
}

class Service2 extends Service {
  execute() {
    console.log("Executing Service2");
  }
}

class ServiceLocator {
  static services = new Map();

  static getService(name) {
    let service = ServiceLocator.services.get(name);

    if (!service) {
      switch (name) {
        case "Service1":
          service = new Service1();
          break;
        case "Service2":
          service = new Service2();
          break;
        default:
          throw new Error("Service not found");
      }
      ServiceLocator.services.set(name, service);
    }

    return service;
  }
}

// Usage
const service = ServiceLocator.getService("Service1");
service.execute();
```

In this example, the ServiceLocator class is a central location for retrieving services. The getService method first checks if the requested service has already been created and stored in the services map. If it has, it returns the cached instance. If not, it creates a new instance of the service and caches it for future use. The Service class acts as an abstract base class for concrete service implementations, such as Service1 and Service2.

Another typical implementation is to use a container framework, such as Spring or Guice, which provides a centralized location for registering and retrieving services and managing their lifecycle.

In conclusion, the Service Locator pattern is a valuable design pattern for managing dependencies between objects in a clean and organized way. By decoupling the creation of services from their use, the Service Locator pattern makes it easier to change the implementation of a service without affecting client objects. It will also improve the performance of an application.

Thank you for taking the time to read this article on the Service Locator design pattern. I hope you found it informative and helpful. If you’d like to stay updated on my future writings, be sure to follow me on Medium. Your support and encouragement in the form of claps would mean a lot to me, and help me to continue producing quality content. Thank you again, and I look forward to connecting with you soon



#### Written on: 2021-10-10 08:00:00
#### Written by: "Evan O'Keeffe"
