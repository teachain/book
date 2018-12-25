<template>
    <Card>
        <Form>
            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('device')" style="width: 100%">添加挖掘机</Button>
            </FormItem>

            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('model')" style="width: 100%">添加型号</Button>
            </FormItem>

            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('driver')" style="width: 100%">添加司机</Button>
            </FormItem>

            <FormItem label="">
                <Button type="primary" v-on:click="linkTo('record')" style="width: 100%">挖掘机出勤记录</Button>
            </FormItem>
           
            <FormItem>
                <Button type="primary" v-on:click="linkTo('maintain')" style="width: 100%">挖掘机保养记录</Button>
            </FormItem>
        </Form>
    </Card>
</template>
<script>
export default {
    name: 'home',
    methods: {
        onSave(name) {
           this.$router.push(name);
        }
    }
}
</script>