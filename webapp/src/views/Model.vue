<template>
    <Card>
        <Form :model="formItem">
            <FormItem label="挖掘机型号" :label-width="80">
                <Input v-model="formItem.name" placeholder="请输入挖掘机型号..." />
            </FormItem>
            <FormItem :label-width="0">
                <Button type="primary" v-on:click="onSave">保存</Button>
            </FormItem>
        </Form>
    </Card>
</template>
<script>
export default {
    name: 'home',
    data() {
        return {
            formItem: {
                name: "",
            }
        }
    },
    methods: {
        onSave() {
            if(this.formItem.name==""){
                this.$Message.error('挖掘机型号不能为空');
                return;
            }
            this.$http.post('/model/add',{name:this.formItem.name}).then(()=>{
                 
            });
            
        }
    }
}
</script>