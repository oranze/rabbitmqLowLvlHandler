package rabbitmqLowLvlHandler

/*
    Generate unique Queue name depend from Environment
*/
func QueueNameGenerate( name *string ) ( resultName string ){
    if Settings.env == "local"{
        resultName = *name +"."+ Settings.identity
    } else{
        resultName = *name +"."+ Settings.env
    }

    return resultName
}