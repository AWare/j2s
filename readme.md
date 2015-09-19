# JSON to Struct (Work in Progress)

This package will generate go code for a struct which matches the structure of types that json.Unmarshal builds when handed a map[string]interface{} as it's unmarshalling type.

    var msi map[string]interface{}
    json.Unmarshal(someJSON, &msi)

The current code quality is: _half past one in the morning_

The current test suite quality is: ___sharp intake of breath over teeth___
