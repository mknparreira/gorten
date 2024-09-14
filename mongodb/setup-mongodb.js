db = db.getSiblingDB('marketplace');

db.createCollection('users', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['userId', 'name', 'email', 'password'],
      properties: {
        userId: {
          bsonType: 'string',
          description: 'must be an string and is required'
        },
        name: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        email: {
          bsonType: 'string',
          pattern: '^.+@.+\..+$',
          description: 'must be a string matching email format and is required'
        },
        password: {
          bsonType: 'string',
          description: 'must be a string and is required'
        }
      }
    }
  }
});

db.createCollection('companies', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['companyId', 'name', 'email'],
      properties: {
        companyId: {
          bsonType: 'string',
          description: 'must be an string and is required'
        },
        name: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        email: {
          bsonType: 'string',
          pattern: '^.+@.+\..+$',
          description: 'must be a string matching email format and is required'
        }
      }
    }
  }
});

db.createCollection('categories', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['categoryId', 'name'],
      properties: {
        categoryId: {
          bsonType: 'string',
          description: 'must be an string and is required'
        },
        name: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        description: {
          bsonType: 'string',
          description: 'must be a string'
        }
      }
    }
  }
});

db.createCollection('products', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['productId', 'name', 'description', 'price'],
      properties: {
        productId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        name: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        description: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        price: {
          bsonType: 'decimal',
          description: 'must be a decimal and is required'
        },
        categoryId: {
          bsonType: 'int',
          description: 'must be an integer representing a category'
        },
        companyId: {
          bsonType: 'int',
          description: 'must be an integer representing a company'
        }
      }
    }
  }
});

db.createCollection('orders', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['orderId', 'userId', 'orderDate', 'totalAmount'],
      properties: {
        orderId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        userId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        orderDate: {
          bsonType: 'date',
          description: 'must be a date and is required'
        },
        totalAmount: {
          bsonType: 'decimal',
          description: 'must be a decimal and is required'
        }
      }
    }
  }
});

db.createCollection('orderItems', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['orderItemId', 'orderId', 'productId', 'quantity', 'price'],
      properties: {
        orderItemId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        orderId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        productId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        quantity: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        price: {
          bsonType: 'decimal',
          description: 'must be a decimal and is required'
        }
      }
    }
  }
});

db.createCollection('payments', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['paymentId', 'orderId', 'paymentDate', 'paymentAmount', 'paymentMethod', 'paymentStatus'],
      properties: {
        paymentId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        orderId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        paymentDate: {
          bsonType: 'date',
          description: 'must be a date and is required'
        },
        paymentAmount: {
          bsonType: 'decimal',
          description: 'must be a decimal and is required'
        },
        paymentMethod: {
          bsonType: 'string',
          description: 'must be a string and is required'
        },
        paymentStatus: {
          bsonType: 'string',
          description: 'must be a string and is required'
        }
      }
    }
  }
});

db.createCollection('shoppingCart', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['cartId', 'userId', 'productId', 'quantity'],
      properties: {
        cartId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        userId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        productId: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        },
        quantity: {
          bsonType: 'int',
          description: 'must be an integer and is required'
        }
      }
    }
  }
});
