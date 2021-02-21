import React from 'react'
import { StyleSheet, Text, View } from 'react-native'

type CodeScreenProps = {
    uid: string
}

export default function CodeScreen({ uid }: CodeScreenProps) {
    return (
        <View>
            <Text>{uid}</Text>
        </View>
    )
}

const styles = StyleSheet.create({})
