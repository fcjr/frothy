import { StatusBar } from 'expo-status-bar'
import React from 'react'
import { View } from 'react-native'
import MainScreen from './screens/MainScreen'
import { NavigationContainer } from '@react-navigation/native'
import { createStackNavigator, StackNavigationOptions, Header, StackHeaderProps } from '@react-navigation/stack'
import FrothyGradient from './components/FrothyGradient'

const GradientHeader = (props: StackHeaderProps) => (
  <View style={{ backgroundColor: '#eee' }}>
      <FrothyGradient>
        <Header {...props} />
      </FrothyGradient>
    </View>
  )

const globalScreenOptions: StackNavigationOptions = {
  header: props => <GradientHeader {...props} />,
  headerStyle: {
    backgroundColor: 'transparent',
  },
  headerTitleStyle: { color: 'white' },
  headerTintColor: 'white'
}

export type RootStackParamList = {
  Main: {}
}

const Stack = createStackNavigator<RootStackParamList>()

export default function App() {
  return (
      <NavigationContainer>
        <StatusBar style='auto' />
        <Stack.Navigator screenOptions={globalScreenOptions}>
          <Stack.Screen name='Main' component={MainScreen} />
        </Stack.Navigator>
      </NavigationContainer>
  )
}