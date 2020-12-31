​		Filter4Go是一款基于GO语言开发的，兼容SQL语言Where子句语法的数据过滤的工具。拥有SQL语言基础的使用者，可以通过编写类SQL的语法规则，对GO环境中的数据进行筛选匹配。例如类似Kubernetes中的各种selector(选择器)，通过配置的选取规则筛选不同资源的应用场景，就是Filter4Go的应用场景之一。另外，对于有些数据量不大，但实时性要求较高的，结构化数据处理场景，Filter4Go也能胜任。如带有数据推送服务的业务中，允许客户自定义数据推送的条件，满足客户的个性化数据需求。

​		Filter4Go是开源项目MOQL(Java)的GO语言版实现，不过它只是实现了MOQL的一个功能子集，即只实现了操作数(Operand)和过滤器(Filter)部分，没有实现筛选器(Selector)与转换器(Translator)两部分。Filter4Go沿用了MOQL的设计思路，独立出了操作数(Operand)部分，而没有将其合并到过滤器(Filter)内部。有Antlr知识基础的人可以看到，filter/filter.g4和oper/oper.g4两个文件完全是可以合并到一起的。但设计上将二者分开，主要是想带入两个子能力，一是定义与执行分离，filter.g4能够解析出一个语法树，这个语法树可以比较容易的与界面配合。界面开发者可以采用树结构做出一个支持Where子句语法的配置界面，并将其转换为Filter4Go语法树的形式，然后再由Filter4Go将其翻译为可执行的Filter操作数，对数据进行过滤。Filter4Go在语法树能力上的支持不如MOQL完善，缺乏从GO对象结构的语法树到类似XML或Json等方便界面操作的格式的转换，如果有需要，可以在后续的版本中进行相应的完善；二是基础操作能力释放，Operand的有效抽离，使得一些如四则混合运算类的配置与使用，可以不依赖于Filter而独立存在，如：希望做一个简单计算器功能就可以直接使用Operand。

​		尽管GO语言并不建议开发者使用它的反射机制，但因为这个工具希望能够更便利，更通用的提供过滤功能，所以不得不使用了相当一部分反射功能。而从希望达到的效果看，在该工具需要的反射支持能力上，GO与Java没有区别。整体而言Filter4Go代码量不大，是很容易看懂的。

Filter4Go四个最重要的接口

- ParseFilterMetadata

​		ParseFilterMetadata(filter string) metadata.OperatableMetadata，将给定的过滤条件解析为语法树结构。

- BuildFilter

​		BuildFilter(metadata metadata.OperatableMetadata) (Operand, error)，基于语法树结构构造为可执行的操作数

- ParseFilter

​		ParseFilter(filter string) (Operand, error)，将给定的过滤条件构造为可执行的操作数

- ParseOperand

​		ParseOperand(operand string) Operand，将给定的表达式解析为可执行的操作数

​		ParseFilter和ParseOperand两个接口都是将一个给定的配置构造为可执行的操作数，那么二者可接收的配置有何不同呢？其实，按照Where子句的语法来说，ParseFilter可接受的语法配置完全兼容ParseOperand的语法配置，即ParseOperand的语法配置是ParseFilter的子集。

​		ParseOperand可接受的语法配置包括：常量(整数、浮点数、字符串、布尔四类)、变量、算数运算符(加、减、乘、除、模)、位运算符(按位与、按位或、异或、非、左移、右移)、结构成员(属性成员、函数成员)、数组、Map、函数。如：“*sprintf('%s:%d', ary[1] ,map['id'])*”，该表达式就是一个可被接受的配置表达式，其中，数组用下标访问，map用key值访问。另外，Filter4Go目前只实现了一个函数，就是sprintf函数，后续会继续扩展支持。如果，使用者需要扩展自己的函数，可以参见functionfactory.go文件，调用FuncFactory的registCreateFunction方法注册自己的函数操作数构造函数。

​		ParseFilter可接受的语法配置除ParseOperand之外，还包括关系运算符(等于、小于、大于、大于等于、小于等于、不等于、in、is、between、like)、逻辑运算符(与、或、非)。如：“bean.Ary[0] = 'A1' and (bean.Id < 50 or bean.Id > 80)”是ParseFilter可接受的过滤条件配置。

​		ParseFilter与ParseOperand除可接受的语法不同外，二者构造返回的Operand的用法也不相同。返回的Operand实际为一个接口，其包括两个重要的执行接口，分别是：

`BooleanOperate(entityMap map[string]interface{}) (bool, error)`

`Operate(entityMap map[string]interface{}) (interface{}, error)`

对ParseFilter返回的Operand，调用其BooleanOperate方法，判断数据是否满足条件；对ParseOperand返回的Operand，调用其Operate方法，对数据进行计算并返回计算结果。两个方法都只接收map类型的参数，map的key必须是字符串，但其值可以是常数、结构体、数组、map等类型。

​		关于Filter4Go的用法示例，可以参看代码工程的filter/filter_test.go和oper/oper_test.go。对于文中解释不详的可以参看[moql](https://github.com/colorknight/moql)。至于如何使用它，可以有更广泛的思考。