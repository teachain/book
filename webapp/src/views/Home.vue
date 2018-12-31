<template>
    <Card>
        <Form>
            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('/device/add')" style="width: 100%">添加挖掘机</Button>
            </FormItem>

            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('/model/add')" style="width: 100%">添加型号</Button>
            </FormItem>

            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('/driver/add')" style="width: 100%">添加司机</Button>
            </FormItem>

            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('/record/add')" style="width: 100%">添加挖掘机出勤记录</Button>
            </FormItem>
            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('/record/list')" style="width: 100%">查看挖掘机出勤记录</Button>
            </FormItem>
           
            <FormItem>
                <Button type="primary" v-on:click="linkTo('/maintain/add')" style="width: 100%">添加挖掘机保养记录</Button>
            </FormItem>
                        <FormItem>
                <Button type="primary" v-on:click="linkTo('/maintain/list')" style="width: 100%">查看挖掘机保养记录</Button>
            </FormItem>
             <FormItem>
                <Button type="primary" v-on:click="linkTo('/gas/add')" style="width: 100%">添加加油记录</Button>
            </FormItem>
             <FormItem>
                <Button type="primary" v-on:click="linkTo('/gas/list')" style="width: 100%">查看加油记录</Button>
            </FormItem>
        </Form>
    </Card>
</template>
<script>
export default {
    methods: {
        linkTo(name) {
           this.$router.push({path:name});
        }
    }
}
</script>