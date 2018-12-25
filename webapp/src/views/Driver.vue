<template>
    <Card>
        <Form :model="formItem">
            <FormItem label="司机名字" :label-width="80">
                <Input v-model="formItem.name" placeholder="请输入司机名字..." />
            </FormItem>
             <FormItem label="手机号码" :label-width="80">
                <Input v-model="formItem.name" placeholder="请输入手机号码..." />
            </FormItem>
            <FormItem :label-width="0">
                <Button type="primary" v-on:click="onSave">保存</Button>
            </FormItem>
        </Form>
    </Card>
</template>
<script>
export default {
    name: 'driver',
    data() {
        return {
            formItem: {
                name: "",
                telephone:"",
            }
        }
    },
    methods: {
        onSave() {
            if(this.formItem.name==""){
                this.$Message.error('司机名字不能为空');
                return;
            }
            if(this.formItem.telephone==""){
                this.$Message.error('司机手机号码不能为空');
                return;
            }
            this.$http.post('/driver/add',{name:this.formItem.name}).then(()=>{
                 
            });
            
        }
    }
}
</script>